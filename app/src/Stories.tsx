import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { RootState } from './store';
import { Story } from './reducers/stories';
import { Container, Grid, Header } from 'semantic-ui-react';
import StoryList from './StoryList';
import StoryView from './StoryView';
import { RootAction } from './actions';
import { listStoriesInit } from './actions/stories';

type StoriesProps = {
  stories: Story[],
  loading: boolean,
  error: Error | null,
  dispatch: Dispatch<RootAction>
};

class Stories extends React.Component<StoriesProps, {}> {

  componentDidMount() {
    this.props.dispatch(listStoriesInit());
  }

  render() {
    return (
      <Container style={{marginTop: '3em'}}>
        <Header as="h1" dividing={true}>Hacker News with gRPC-Web</Header>

        <Grid columns={2} stackable={true} divided={'vertically'}>
          <Grid.Column width={3}>
            <StoryList stories={this.props.stories}/>
          </Grid.Column>

          <Grid.Column width={13} stretched={true}>

            <Header as="h2">Example body text</Header>
            <StoryView/>
          </Grid.Column>
        </Grid>

      </Container>
    );
  }

}

export default connect((state: RootState) => ({
  stories: Array.from(state.stories.stories.values()),
  loading: state.stories.loading,
  error: state.stories.error,
}))(Stories);
